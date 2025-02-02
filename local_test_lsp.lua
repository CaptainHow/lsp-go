local client = vim.lsp.start_client({
	name = "go-lsp",
	cmd = { "D:/AHMED/Projects/Golang/LSP/lsp-go/main" },
	on_attach = function() end,
})

if not client then
	vim.notify("Hey, client issue")
	return
end

vim.api.nvim_create_autocmd("FileType", {
	pattern = "markdown",
	callback = function()
		vim.lsp.buf_attach_client(0, client)
	end,
})
