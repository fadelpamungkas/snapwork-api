let s:cpo_save=&cpo
set cpo&vim
inoremap <silent> <Up> <Cmd>call v:lua.cmp.utils.keymap.set_map(12)
inoremap <silent> <Down> <Cmd>call v:lua.cmp.utils.keymap.set_map(7)
inoremap <silent> <S-Tab> <Cmd>call v:lua.cmp.utils.keymap.set_map(4)
cnoremap <silent> <Plug>(TelescopeFuzzyCommandSearch) e "lua require('telescope.builtin').command_history { default_text = [=[" . escape(getcmdline(), '"') . "]=] }"
noremap! <silent> <Plug>luasnip-expand-repeat <Cmd>lua require'luasnip'.expand_repeat()
noremap! <silent> <Plug>luasnip-delete-check <Cmd>lua require'luasnip'.unlink_current_if_deleted()
inoremap <silent> <Plug>luasnip-jump-prev <Cmd>lua require'luasnip'.jump(-1)
inoremap <silent> <Plug>luasnip-jump-next <Cmd>lua require'luasnip'.jump(1)
inoremap <silent> <Plug>luasnip-prev-choice <Cmd>lua require'luasnip'.change_choice(-1)
inoremap <silent> <Plug>luasnip-next-choice <Cmd>lua require'luasnip'.change_choice(1)
inoremap <silent> <Plug>luasnip-expand-snippet <Cmd>lua require'luasnip'.expand()
inoremap <silent> <Plug>luasnip-expand-or-jump <Cmd>lua require'luasnip'.expand_or_jump()
nnoremap  h
snoremap <silent> 	 <Cmd>call v:lua.cmp.utils.keymap.set_map(5)
nnoremap <NL> j
nnoremap  k
nnoremap  l
nnoremap <silent>  <Cmd>execute v:count . "ToggleTerm"
nnoremap <silent>  dq <Cmd>lua vim.diagnostic.setloclist()
nnoremap <silent>  d] <Cmd>lua vim.diagnostic.goto_next()
nnoremap <silent>  d[ <Cmd>lua vim.diagnostic.goto_prev()
nnoremap <silent>  dh <Cmd>lua vim.diagnostic.disable()
nnoremap <silent>  ds <Cmd>lua vim.diagnostic.enable()
nnoremap <silent>  dd <Cmd>lua vim.diagnostic.open_float()
nnoremap <silent>  ' <Cmd>Telescope registers
nnoremap <silent>  o <Cmd>Telescope oldfiles
nnoremap <silent>  b <Cmd>Telescope buffers
nnoremap <silent>  g <Cmd>Telescope live_grep
nnoremap <silent>  f <Cmd>Telescope find_files find_command=rg,--ignore,--files
nnoremap  sh <Cmd>split
nnoremap  sv <Cmd>vsplit
nnoremap  - <Cmd>resize -2
nnoremap  = <Cmd>resize +2
nnoremap  , <Cmd>:vertical resize -2
nnoremap  . <Cmd>:vertical resize +2
nmap <silent>   <Nop>
omap <silent> % <Plug>(MatchitOperationForward)
xmap <silent> % <Plug>(MatchitVisualForward)
nmap <silent> % <Plug>(MatchitNormalForward)
vnoremap < <gv
vnoremap > >gv
xmap S <Plug>VSurround
nnoremap Y y$
omap <silent> [% <Plug>(MatchitOperationMultiBackward)
xmap <silent> [% <Plug>(MatchitVisualMultiBackward)
nmap <silent> [% <Plug>(MatchitNormalMultiBackward)
omap <silent> ]% <Plug>(MatchitOperationMultiForward)
xmap <silent> ]% <Plug>(MatchitVisualMultiForward)
nmap <silent> ]% <Plug>(MatchitNormalMultiForward)
xmap a% <Plug>(MatchitVisualTextObject)
nmap cS <Plug>CSurround
nmap cs <Plug>Csurround
nmap ds <Plug>Dsurround
xmap gS <Plug>VgSurround
nmap gcu <Plug>Commentary<Plug>Commentary
nmap gcc <Plug>CommentaryLine
omap gc <Plug>Commentary
nmap gc <Plug>Commentary
xmap gc <Plug>Commentary
xmap gx <Plug>NetrwBrowseXVis
nmap gx <Plug>NetrwBrowseX
omap <silent> g% <Plug>(MatchitOperationBackward)
xmap <silent> g% <Plug>(MatchitVisualBackward)
nmap <silent> g% <Plug>(MatchitNormalBackward)
nmap ySS <Plug>YSsurround
nmap ySs <Plug>YSsurround
nmap yss <Plug>Yssurround
nmap yS <Plug>YSurround
nmap ys <Plug>Ysurround
snoremap <silent> <S-Tab> <Cmd>call v:lua.cmp.utils.keymap.set_map(3)
nnoremap <silent> <Plug>SurroundRepeat .
nmap <silent> <Plug>CommentaryUndo :echoerr "Change your <Plug>CommentaryUndo map to <Plug>Commentary<Plug>Commentary"
nnoremap <Plug>PlenaryTestFile :lua require('plenary.test_harness').test_directory(vim.fn.expand("%:p"))
snoremap <silent> <Plug>luasnip-jump-prev <Cmd>lua require'luasnip'.jump(-1)
snoremap <silent> <Plug>luasnip-jump-next <Cmd>lua require'luasnip'.jump(1)
snoremap <silent> <Plug>luasnip-prev-choice <Cmd>lua require'luasnip'.change_choice(-1)
snoremap <silent> <Plug>luasnip-next-choice <Cmd>lua require'luasnip'.change_choice(1)
snoremap <silent> <Plug>luasnip-expand-snippet <Cmd>lua require'luasnip'.expand()
snoremap <silent> <Plug>luasnip-expand-or-jump <Cmd>lua require'luasnip'.expand_or_jump()
noremap <silent> <Plug>luasnip-expand-repeat <Cmd>lua require'luasnip'.expand_repeat()
noremap <silent> <Plug>luasnip-delete-check <Cmd>lua require'luasnip'.unlink_current_if_deleted()
xnoremap <silent> <Plug>NetrwBrowseXVis :call netrw#BrowseXVis()
nnoremap <silent> <Plug>NetrwBrowseX :call netrw#BrowseX(netrw#GX(),netrw#CheckIfRemote(netrw#GX()))
xmap <silent> <Plug>(MatchitVisualTextObject) <Plug>(MatchitVisualMultiBackward)o<Plug>(MatchitVisualMultiForward)
onoremap <silent> <Plug>(MatchitOperationMultiForward) :call matchit#MultiMatch("W",  "o")
onoremap <silent> <Plug>(MatchitOperationMultiBackward) :call matchit#MultiMatch("bW", "o")
xnoremap <silent> <Plug>(MatchitVisualMultiForward) :call matchit#MultiMatch("W",  "n")m'gv``
xnoremap <silent> <Plug>(MatchitVisualMultiBackward) :call matchit#MultiMatch("bW", "n")m'gv``
nnoremap <silent> <Plug>(MatchitNormalMultiForward) :call matchit#MultiMatch("W",  "n")
nnoremap <silent> <Plug>(MatchitNormalMultiBackward) :call matchit#MultiMatch("bW", "n")
onoremap <silent> <Plug>(MatchitOperationBackward) :call matchit#Match_wrapper('',0,'o')
onoremap <silent> <Plug>(MatchitOperationForward) :call matchit#Match_wrapper('',1,'o')
xnoremap <silent> <Plug>(MatchitVisualBackward) :call matchit#Match_wrapper('',0,'v')m'gv``
xnoremap <silent> <Plug>(MatchitVisualForward) :call matchit#Match_wrapper('',1,'v'):if col("''") != col("$") | exe ":normal! m'" | endifgv``
nnoremap <silent> <Plug>(MatchitNormalBackward) :call matchit#Match_wrapper('',0,'n')
nnoremap <silent> <Plug>(MatchitNormalForward) :call matchit#Match_wrapper('',1,'n')
inoremap <silent>  <Cmd>call v:lua.cmp.utils.keymap.set_map(1)
inoremap <silent>  <Cmd>call v:lua.cmp.utils.keymap.set_map(10)
inoremap <silent>  <Cmd>call v:lua.cmp.utils.keymap.set_map(11)
imap S <Plug>ISurround
imap s <Plug>Isurround
inoremap <silent> 	 <Cmd>call v:lua.cmp.utils.keymap.set_map(6)
inoremap <silent>  <Cmd>call v:lua.cmp.utils.keymap.set_map(8)
inoremap <silent>  <Cmd>call v:lua.cmp.utils.keymap.set_map(2)
inoremap <silent>  <Cmd>call v:lua.cmp.utils.keymap.set_map(13)
imap  <Plug>Isurround
inoremap  u
inoremap  u
inoremap <silent>  <Cmd>call v:lua.cmp.utils.keymap.set_map(14)
inoremap <silent>  <Cmd>ToggleTerm
inoremap <silent> ;; <Cmd>call v:lua.cmp.utils.keymap.set_map(9)
inoremap jk 
let &cpo=s:cpo_save
unlet s:cpo_save
set clipboard=unnamedplus
set completeopt=menu,menuone,noselect
set expandtab
set nohlsearch
set ignorecase
set indentkeys=0{,0},0),0],:,0#,!^F,o,O,e,<:>,0=},0=)
set mouse=a
set operatorfunc=<SNR>35_go
set runtimepath=~/.config/nvim,/etc/xdg/nvim,~/.local/share/nvim/site,~/.local/share/nvim/site/pack/*/start/*,~/.local/share/nvim/site/pack/packer/start/packer.nvim,/usr/local/share/nvim/site,/usr/share/nvim/site,/opt/homebrew/Cellar/neovim/0.7.0/share/nvim/runtime,/opt/homebrew/Cellar/neovim/0.7.0/share/nvim/runtime/pack/dist/opt/matchit,/opt/homebrew/Cellar/neovim/0.7.0/lib/nvim,~/.local/share/nvim/site/pack/*/start/*/after,/usr/share/nvim/site/after,/usr/local/share/nvim/site/after,~/.local/share/nvim/site/after,/etc/xdg/nvim/after,~/.config/nvim/after
set shiftwidth=2
set showmatch
set smartcase
set softtabstop=2
set splitbelow
set splitright
set statusline=%{%v:lua.require'lualine'.statusline()%}
set tabstop=2
set termguicolors
set undofile
set updatetime=250
set window=57
" vim: set ft=vim :
