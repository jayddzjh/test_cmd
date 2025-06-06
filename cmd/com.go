package cmd

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var commitMessage string // 用于存储 -m 标志的值

var commitCmd = &cobra.Command{
    Use:   "commit",
    Short: "记录对仓库的更改 (等同于 git commit)",
    Long: `此命令将暂存区的更改保存到本地仓库。
目前仅支持通过 -m 标志提供提交信息。`,
    RunE: func(cmd *cobra.Command, args []string) error {
        if commitMessage == "" {
            // Cobra 的 MarkFlagRequired 应该会处理这个，但作为双重检查
            return fmt.Errorf("提交信息是必需的，请使用 -m \"message\"")
        }

        // 构建要执行的 git 命令
        gitArgs := []string{"commit", "-m", commitMessage}

        gitCmd := exec.Command("git", gitArgs...)
        gitCmd.Stdout = os.Stdout
        gitCmd.Stderr = os.Stderr

        fmt.Printf("执行: git %s\n", stringArrayToString(gitArgs))

        err := gitCmd.Run()
        if err != nil {
            return fmt.Errorf("执行 'git commit' 失败: %w", err)
        }
        fmt.Println("提交成功。")
        return nil
    },
}

func init() {
    rootCmd.AddCommand(commitCmd) // 将 commitCmd 添加为 rootCmd 的子命令

    // 为 commitCmd 添加一个标志
    // StringVarP 允许我们定义一个字符串类型的标志，并将其值绑定到 commitMessage 变量
    // "message" 是长标志名 (--message)
    // "m" 是短标志名 (-m)
    // "" 是默认值
    // "提交信息 (必需)" 是帮助文本
    commitCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "提交信息 (必需)")

    // 将 -m 标志标记为必需
    // 如果用户没有提供这个标志，Cobra 会自动报错
    if err := commitCmd.MarkFlagRequired("message"); err != nil {
        // 在 init 阶段处理错误通常是 panic，因为这是开发时的问题
        panic(fmt.Sprintf("标记 'message' 标志为必需时出错: %v", err))
    }
}
