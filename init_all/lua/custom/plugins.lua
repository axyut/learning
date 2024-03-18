local plugins = {
  {
    "vim-crystal/vim-crystal",
    ft = "crystal",
    config = function (_)
      vim.g.crystal_auto_format = 1
    end
  },{
    "neovim/nvim-lspconfig",
    config = function ()
      require "plugins.configs.lspconfig"
      require "custom.configs.lspconfig"
    end,
  },
  {
    "williamboman/mason.nvim",
    opts = {
      ensure_installed = {
        "rust_analyzer",
        "gopls",
        "golines",
        "goimports",
        "gofumpt"
      },
    },
  },
  {
    "jose-elias-alvarez/null-ls.nvim",
    ft = "go",
    opts = function ()
      return require "custom.configs.null-ls"
    end
  },
  {
    "olexsmir/gopher.nvim",
    ft = "go",
    config = function (_, opts)
      require("gopher").setup(opts)
    end,
    build = function ()
      vim.cmd [[silent! GoInstallDeps]]
    end
  }
}
return plugins
