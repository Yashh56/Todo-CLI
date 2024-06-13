package cmd

func init() {
	rootCmd.AddCommand(TodoCmd)
	TodoCmd.Flags().BoolVarP(&AddFlag, "add", "a", false, "Add a new todo item with the specified title")
	TodoCmd.Flags().BoolVarP(&DeleteFlag, "remove", "r", false, "Delete the todo item with the specified title")
	TodoCmd.Flags().BoolVarP(&UpdateFlag, "update", "u", false, "Update the todo completed or not")
	TodoCmd.Flags().BoolVarP(&RemoveAllFlag, "removeAll", "x", false, "Update the todo completed or not")
	rootCmd.AddCommand(GetAllCmd)
	rootCmd.AddCommand(GetAllCompletedCmd)
	rootCmd.AddCommand(GetPendingCmd)
	rootCmd.AddCommand(DeleteAllCmd)
}
