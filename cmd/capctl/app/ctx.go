package app

import (
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/ctx"
	"github.com/spf13/cobra"
	"strings"
)

const (
	ctxHelp = `切換不同的 captain-kube 環境

Context 抽象化了 flags, 你可以視一個 context 為一組環境設定
將配置好的 context 啟用後, 會在執行任何 command 前被載入, command 使用的順序為:

	1. 當前 command 執行時所傳入的 flag 
	2. 啟用中的 context-flags 
	3. 環境變數的設定

'{{.}} ctx' 指令可以快速的在不同 context 之間切換

	$ {{.}} ctx         : 互動式的快速切換 context
	$ {{.}} ctx <NAME>  : 切換 context 到 <NAME>
	$ {{.}} ctx -       : 切換到前一個 context
	$ {{.}} ctx --off   : 清空當前的 context

傳入 '--ls' 可以列出所有 context 及其 context-flags
配合 '--width' 可以指定顯示的字數 (預設100), '--width 0' 為不限長度, 即完整顯示

	$ {{.}} ctx --ls
	$ {{.}} ctx --ls --width 0

傳入 '--add' 可以新增 context, 結合 '--force' 可以強制複寫已存在的 context
使用上需先接著一組 double-dash (--), 之後再給予1到數個 context-flags

	$ {{.}} ctx -a <NAME> -- <CONTEXT_FLAGS...>
	$ {{.}} ctx -a local -f -- -e localhost --endpoint-port 30051  

傳入 '--delete' 或 '--rename' 可以刪除或重新命名 context:

	$ {{.}} ctx -d <NAME>             : 刪除 context <NAME>
	$ {{.}} ctx -d .                  : 刪除當前的 context
	$ {{.}} ctx -r <NAME>=<NEW_NAME>  : 重新命名 <NAME> 成 <NEW_NAME>
	$ {{.}} ctx -r .=<NEW_NAME>       : 重新命名當前的 context name 成 <NEW_NAME>

可用的 context-flags 包含了:

%s
`
)

func formatCtxHelp() string {
	help := usage(ctxHelp)
	flags, _ := ctx.FlagsString()
	return fmt.Sprintf(help, flags)
}

type ctxCmd struct {
	width  uint
	add    string
	rename string
	force  bool
	ls     bool
	off    bool
	delete []string
	args   []string
}

func newCtxCmd() *cobra.Command {
	c := ctxCmd{}

	cmd := &cobra.Command{
		Use:   "ctx",
		Short: "切換不同的 captain-kube 環境",
		Long:  formatCtxHelp(),
		Args: func(cmd *cobra.Command, args []string) error {
			if ctxs == ctx.PlainContexts {
				return ctx.ErrMountVolumeNotExist
			}
			if len(c.add) > 0 && len(args) == 0 {
				return fmt.Errorf("requires at least 1 argument to add context")
			}
			if len(c.delete) > 0 && len(args) > 0 {
				return fmt.Errorf("delete context does not accpet arguments")
			}
			if c.off && len(args) > 0 {
				return fmt.Errorf("switch off context does not accpet arguments")
			}
			if len(c.rename) > 0 {
				if len(args) > 0 {
					return fmt.Errorf("rename context does not accpet arguments")
				}
				if !strings.Contains(c.rename, "=") {
					return fmt.Errorf("requires 1 equal sign (=) to rename context, e.g. <NAME>=<NEW_NAME>")
				}
			}
			if len(c.add) == 0 && len(c.delete) == 0 && len(args) > 1 {
				return fmt.Errorf("list/switch context only accpet max 1 argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c.args = args
			return c.run()
		},
	}

	f := cmd.Flags()
	f.StringVarP(&c.add, "add", "a", "", "add context <NAME> with <CONTEXT_FLAGS...>")
	f.BoolVarP(&c.force, "force", "f", false, "force to delete context if exist")
	f.StringArrayVarP(&c.delete, "delete", "d", []string{}, "delete context <NAME> ('.' for current-context)")
	f.StringVarP(&c.rename, "rename", "r", "", "rename context <NAME> to <NEW_NAME>")
	f.BoolVar(&c.ls, "ls", false, "list contexts")
	f.BoolVar(&c.off, "off", false, "switch off the context")
	f.UintVar(&c.width, "width", 100, "maximum allowed width for listing context args")

	return cmd
}

func (c *ctxCmd) run() error {
	if c.off {
		return ctxs.SwitchOff()
	}
	if len(c.add) > 0 {
		return ctxs.Add(c.add, c.args, c.force)
	}
	if len(c.delete) > 0 {
		return ctxs.Delete(c.delete...)
	}
	if len(c.rename) > 0 {
		r := strings.Split(c.rename, "=")
		return ctxs.Rename(r[0], r[1])
	}
	if len(c.args) > 0 {
		return ctxs.Switch(c.args[0])
	}

	names := ctxs.GetSortedNames()

	if c.ls {
		table := uitable.New()
		table.AddRow("CONTEXT", "FLAGS")
		table.MaxColWidth = c.width
		for _, name := range names {
			args := ctxs.Contexts[name]
			prefix := " "
			if name == ctxs.Active {
				prefix = ">"
			} else if name == ctxs.Previous {
				prefix = "-"
			}
			table.AddRow(fmt.Sprintf("%s %s", prefix, name), strings.Join(args, " "))
		}
		logrus.Println(table)
		return nil
	}

	var items []string
	for _, ctx := range names {
		prefix := " "
		if ctx == ctxs.Active {
			prefix = ">"
		}
		// 都要用過濾的了, 好像不需要再顯示前一個了
		//else if ctx == ctxs.Previous {
		//	prefix = "-"
		//}
		items = append(items, fmt.Sprintf("%s %s", prefix, ctx))
	}

	prompt := promptui.Select{
		Label: "Select Context",
		Items: items,
		Searcher: func(input string, index int) bool {
			ctx := items[index]
			name := strings.Replace(strings.ToLower(ctx), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return err
	}
	return ctxs.Switch(result[2:])
}
