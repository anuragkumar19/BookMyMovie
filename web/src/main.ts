import Lara from "@primevue/themes/lara"
import PrimeVue from "primevue/config"
import { createApp } from "vue"
import App from "./App.vue"
import Router from "./router"
import { ConnectTransport } from "./services/connect-transport"

const app = createApp(App)

app.use(PrimeVue, {
	theme: {
		preset: Lara,
	},
})
app.use(ConnectTransport)
app.use(Router)

app.mount("#app")
