package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"log"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "This run the fiber development server",
	Long:  `This run the fiber development server`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
func startServer() {
	app := fx.New(
		fx.Provide(initFiberApp),
		fx.Invoke(startFiberServer, loadEnv),
	)

	app.Run()
}

func initFiberApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4200",
		AllowCredentials: true,
	}))

	return app
}

func registerGeneratedHandler(
	app *fiber.App,
	// services server.Server,
) {
	//api.RegisterHandlers(app, services)
}

func startFiberServer(app *fiber.App) {
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}

func loadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	//if value, exists := os.LookupEnv("DB_DSN"); !exists || value == "" {
	//	if err := os.Setenv("DB_DSN", "local.db"); err != nil {
	//		return err
	//	}
	//}

	return nil
}
