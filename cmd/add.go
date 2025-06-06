package cmd

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add <file1> [file2...]", // 命令使用方式
    Short: "将文件内容添加到索引 (等同于 git add)",
    Long: `此命令将指定的文件内容添加到 Git 索引中。
这和执行 'git add <file...>' 效果相同。`,
    Args: cobra.MinimumNArgs(1), // 至少需要一个参数（文件名）
    RunE: func(cmd *cobra.Command, args []string) error { // 使用 RunE 来处理错误
        // 构建要执行的 git 命令
        // 第一个参数是 "git"，后续是 "add" 和用户传入的文件名
        gitArgs := append([]string{"add"}, args...)

        // 创建一个 *exec.Cmd 对象
        gitCmd := exec.Command("git", gitArgs...)

        // 将子进程的 stdout 和 stderr 连接到当前进程的 stdout 和 stderr
        // 这样用户就能直接看到 git 命令的输出
        gitCmd.Stdout = os.Stdout
        gitCmd.Stderr = os.Stderr

        fmt.Printf("执行: git %s\n", stringArrayToString(gitArgs)) // 打印将要执行的命令

        // 执行命令
        err := gitCmd.Run()
        if err != nil {
            return fmt.Errorf("执行 'git add' 失败: %w", err)
        }
        fmt.Println("文件已成功添加到暂存区。")
        return nil // 成功时返回 nil
    },
}

// init 函数在包被导入时自动执行
func init() {
    rootCmd.AddCommand(addCmd) // 将 addCmd 添加为 rootCmd 的子命令
}

// 辅助函数，将字符串切片转换为用空格分隔的字符串（用于打印）
func stringArrayToString(arr []string) string {
    var result string
    for i, s := range arr {
        if i > 0 {
            result += " "
        }
        result += s
    }
    return result
}
