package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"runtime"
)

var configDefaults = map[string]interface{}{
	"gomaxprocs": 0,
	"java_host":  "127.0.0.1",
	"java_port":  "8001",
	"host":       "127.0.0.1",
	"port":       "8000",
}

var (
	host     string
	port     string
	javaHost string
	javaPort string
)

func init() {
	rootCmd.Flags().StringVarP(&host, "host", "o", "127.0.0.1", "http host")
	rootCmd.Flags().StringVarP(&port, "port", "d", "8000", "http port")
	rootCmd.Flags().StringVarP(&javaHost, "java_host", "m", "127.0.0.1", "java host")
	rootCmd.Flags().StringVarP(&javaPort, "java_port", "k", "8001", "java port")
}

var rootCmd = &cobra.Command{
	Use: "rtk-backend",
	Run: func(cmd *cobra.Command, args []string) {

		for k, v := range configDefaults {
			viper.SetDefault(k, v)
		}

		bindEnvs := []string{
			"host", "port", "java_host", "java_port",
		}
		for _, env := range bindEnvs {
			err := viper.BindEnv(env)
			if err != nil {
				logrus.Fatalf("error binding env variable: %v", err)
			}
		}

		if os.Getenv("GOMAXPROCS") == "" {
			if viper.IsSet("gomaxprocs") && viper.GetInt("gomaxprocs") > 0 {
				runtime.GOMAXPROCS(viper.GetInt("gomaxprocs"))
			} else {
				runtime.GOMAXPROCS(runtime.NumCPU())
			}
		}

		h := &Handler{}

		r := gin.Default()
		setupRouter(r, h)
		err := r.Run()
		if err != nil {
			logrus.Errorf("error running gin: %v", err)
		}
	},
}

func setupRouter(r *gin.Engine, h *Handler) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("login", h.HandleAuth)
}

func main() {

	err := rootCmd.Execute()
	if err != nil {
		logrus.Errorf("error running root command: %v", err)
	}

}
