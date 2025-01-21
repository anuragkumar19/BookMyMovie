import type { User } from "./services/gen/users/v1/users_pb"
import { ConnectError } from "@connectrpc/connect"
import { computed, ref } from "vue"
import { AuthService } from "./services/gen/auth/v1/auth_connect"
import { UsersService } from "./services/gen/users/v1/users_connect"
import { useService } from "./services/use-service"

export function useAuth() {
	const baseUser = ref<User | null>(null)
	const isInitiated = ref<boolean>(false)
	const isFetching = ref<boolean>(false)

	const accessToken = ref<string | null>(null)
	const accessTokenExpiry = ref<Date | null>(null)

	const authService = useService(AuthService)
	const usersService = useService(UsersService)

	async function fetchUser() {
		if (!isInitiated.value) {
			await getAccessToken()
		}
	}

	async function getAccessToken() {
		if (!isInitiated.value && (accessToken.value === null || accessTokenExpiry.value == null)) {
			return null
		}

		if (accessTokenExpiry.value < new Date(Date.now() + 10 * 1000)) {
			try {
				const res = await authService.refreshAccessToken({}, {
					headers: {
						Authorization: `Bearer ${localStorage.getItem("refresh_token")}`,
					},
				})

				accessToken.value = res.accessToken
				accessTokenExpiry.value = res.accessTokenExpiry!.toDate()
			}
			catch (err) {
				if (err instanceof ConnectError) {
					alert(err.message)
				}
			}
		}
	}

	return {
		user: computed(() => baseUser.value),
		isFetching: computed(() => isFetching.value),
		isLoggedIn: computed(() => baseUser.value !== null),
		accessToken: computed(() => accessToken.value),
		fetchUser,
	}
}
