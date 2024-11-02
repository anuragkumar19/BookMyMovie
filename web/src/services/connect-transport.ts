import type { Transport } from "@connectrpc/connect"
import { createConnectTransport } from "@connectrpc/connect-web"
import { type App, inject, type InjectionKey } from "vue"

const key = Symbol("connect-transport") as InjectionKey<Transport>

export const ConnectTransport = {
	install(app: App) {
		const transport = createConnectTransport({
			baseUrl: import.meta.env.VITE_CONNECTRPC_BASE_URL,
			fetch,
			useBinaryFormat: true,
		})
		app.provide(key, transport)
	},
}

export function useTransport(): Transport {
	const transport = inject(key)
	if (!transport) {
		throw new Error("cannot use transport without registering ConnectTransport plugin in Vue.App")
	}
	return transport
}
