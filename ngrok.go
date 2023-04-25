package ngrok

import(
    "context"
    "log"
    "net/http"
    
    "golang.ngrok.com/ngrok"
    "golang.ngrok.com/ngrok/config"
)

func ngrok() {
    
    if err:= run(context.Background()); err != nil {
        log.Fatal(err)
    }

}

func run(ctx context.Context) error {
    tun, err := ngrok.Listen(ctx,
        config.HTTPEndpoint(),
        ngrok.WithAuthtokenFromEnv(),
    )
    if err != nil {
        return err
    }
    
    log.Println("tunnel created:", tun.URL())
    
    return http.Serve(tun, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "<h1>Hello from ngrok-go.</h1>")
}