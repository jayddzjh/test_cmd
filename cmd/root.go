package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "mygit", //  命令名
    Short: "一个简单的 Git 命令包装器",
    Long: `mygit 是一个使用 Cobra 库构建的 CLI 工具，
它封装了一些常用的 Git 命令，用于学习目的。`,
    // 如果 rootCmd 本身有操作，可以在这里定义 Run 函数
    // Run: func(cmd *cobra.Command, args []string) { },
}

// Execute 函数由 main.main() 调用。它只需要在 rootCmd 上调用一次。
// 这个函数会解析命令和参数，然后执行相应的 Run 函数
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "执行命令时出错: '%s'", err)
        os.Exit(1)
    }
}

// init 函数在包被导入时自动执行
// 我们在这里可以进行一些初始化设置，比如添加全局标志等
func init() {
    //rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
