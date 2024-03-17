/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/brunowang/gokv/internal/engine"
	"github.com/brunowang/gokv/internal/raftmgr"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cpath == "" {
			panic(fmt.Errorf("config path %s not found", cpath))
		}
		if err := raftmgr.Init(cpath); err != nil {
			panic(err)
		}

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			defer func() {
				if e := recover(); e != nil {
					c.JSON(400, gin.H{"message": e})
				}
			}()
			c.Next()
		})
		router.Handle("GET", "/get", func(c *gin.Context) {
			key := c.Query("key")
			value := engine.Get(key)
			c.JSON(200, gin.H{
				"key":   key,
				"value": value,
			})
		})
		router.Handle("POST", "/set", func(c *gin.Context) {
			var req engine.KVPair
			if err := c.BindJSON(&req); err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			future := raftmgr.RaftNode.Apply(req.ToBytes(), 500*time.Millisecond)
			if err := future.Error(); err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
		err := router.Run(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatal("Router.Run error: ", err)
		}
	},
}

var (
	cpath string
	port  uint32
)

func init() {
	rootCmd.AddCommand(httpCmd)

	httpCmd.Flags().StringVarP(&cpath, "config", "c", "", "config path")
	httpCmd.Flags().Uint32VarP(&port, "port", "p", 10080, "http port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
