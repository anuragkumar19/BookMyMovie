import antfu from "@antfu/eslint-config"

export default antfu({
	vue: true,
	typescript: true,
	stylistic: {
		quotes: "double",
		indent: "tab",
		semi: false,
	},
	ignores: ["src/services/gen/**/*"],
})
