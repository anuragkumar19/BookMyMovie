import type { ServiceType } from "@bufbuild/protobuf"

import { createPromiseClient, type PromiseClient } from "@connectrpc/connect"
import { useTransport } from "./connect-transport"

export function useService<T extends ServiceType>(s: T): PromiseClient<T> {
	const transport = useTransport()
	const client = createPromiseClient(s, transport)

	return client
}
