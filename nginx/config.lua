local config = require("lapis.config")

config("development", {
    host = os.getenv("SERVER_PORT_8000_TCP_ADDR"),
    port = os.getenv("SERVER_PORT_8000_TCP_PORT")
})
