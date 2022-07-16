" ???
let mapleader=" "
call plug#begin('~/../Local/nvim/plugged')
" file searching
Plug 'dracula/vim'
Plug 'junegunn/fzf', { 'on': 'FZF' }
Plug 'junegunn/fzf.vim'
Plug 'morhetz/gruvbox'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'preservim/nerdtree'
call plug#end()
"Config Section
if (has("termguicolors"))
 set termguicolors
endif
syntax enable

let g:gruvbox_contrast_dark='hard'
 colorscheme dracula

": set background=dark    " Setting dark mode

set number

let g:NERDTreeShowHidden = 1
let g:NERDTreeMinimalUI i
let g:NERDTreeIgnore = []
let g:NERDTreeStatusline = ''

""""""""""""""""""""""""""""""""""""""""""
" Plugin 'The-NERD-tree'
""""""""""""""""""""""""""""""""""""""""""
let g:NERDChristmasTree=1
let g:NERDTreeAutoCenter=1
let g:NERDTreeChDirMode=1
let g:NERDTreeHighlightCursorline=1
let g:NERDTreeSortOrder=['\.py$', '\.m$', '\.pyc$', '*']
let g:NERDTreeBookmarksFile = $HOME.'/.vim/NERDTreeBookmarks'
let g:NERDTreeIgnore=['\.vim$', '\~$']
let g:NERDTreeWinSize=40
let g:NERDTreeMinimalUI=1
let g:NERDTreeShowHidden=1
let g:NERDTreeShowLineNumbers=1
let g:NERDTreeQuitOnOpen=1
let g:NERDTreeShowFiles=1
let g:NERDTreeWinPos="left"

""""""""""""""""""""""""""""""""""""""""""	
set clipboard=unnamedplus

set t_Co=256
colorscheme wombat256mod
"colorscheme gruvbox

set cursorcolumn                           " highlight the current column
set cursorline                             " highlight current line
set incsearch                              " BUT do highlight as you type you search phrase
set laststatus=2                           " always show the status line
set lazyredraw                             " do not redraw while running macros
set linespace=0                            " don't insert any extra pixel line betweens rows
"set list                                  " we do what to show tabs, to
"ensure we get them out of my files
""set listchars=tab:>-,trail:.,extends:>    " show tabs and trailing
set matchtime=5                            " how many tenths of a second to blink matching brackets for
set hlsearch                               " do not highlight searched for phrases
set nostartofline                          " leave my cursor where it was
set novisualbell                           " don't blink
set noerrorbells                           " don't beep
set number                                 " turn on line numbers
set numberwidth=5                          " We are good up to 99999 lines
set report=0                               " tell us when anything is changed via :...
set ruler                                  " Always show current positions along the bottom
set scrolloff=10                           " Keep 10 lines (top/bottom) for scope
set shortmess=aOstT                        " shortens messages to avoid 'press a key' prompt
set showcmd                                " show the command being typed
set showmatch                              " show matching brackets
set sidescrolloff=10
set statusline+=%#warningmsg#
set statusline+=%*


set encoding=utf-8   

set backspace=2
set smartindent
set autoindent
set copyindent
set cindent
set nobackup
set mouse=a
set noto ttimeout
set autoread
"set completeopt=                          " don't use a pop up menu for completions
set expandtab                              " no real tabs please!
set formatoptions-=cro                     " Automatically insert comment
"leader on return, and let gq format comments
"set ignorecase                            " case insensitive by default
set infercase                              " case inferred by default
set nowrap                                 " do not wrap line
set shiftround                             " when at 3 spaces, and I hit > ... go to 4, not 5
set smartcase                              " if there are caps, go case-sensitive
set shiftwidth=4                           " auto-indent amount when using cindent, >>, << and stuff like that
set shiftround                             " use multiple of shiftwidth when indenting with '<' and '>'
set softtabstop=4                          " when hitting tab or backspace, how many spaces should a tab be (see expandtab)
set tabstop=4                              " real tabs should be 8, and they will show with set list on
set ww=h,l,b                               " move the cursor left/right to move to the previous/next line when the cursor is on the first/last character in the line
set sol                                    " when moving cursor(ex:<C-d>), move the cursor to the start of line
set wmh=0                                  " minium hegith of window
set title                                  " change name of terminal
set pastetoggle=<leader>p                  " toggle vim into past mode
set ttyfast                                " fast keyboard
set timeout timeoutlen=1000 ttimeoutlen=0 " fast keyboard

let g:airline_powerline_fonts=1
" let g:airline_theme='gruvbox'
let g:airline#extensions#tabline#enabled=1
let g:airline_left_sep = '»'
let g:airline_left_sep = '▶'
let g:airline_right_sep = '«'
let g:airline_right_sep = '◀'



" Automaticaly close nvim if NERDTree is only thing left open
autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTree") && b:NERDTree.isTabTree()) | q | endif
" Toggle
nnoremap <silent> <C-b> :NERDTreeToggle<CR>

" open new split panes to right and below
set splitright
set splitbelow
" turn terminal to normal mode with escape
tnoremap <Esc> <C-\><C-n>
" start terminal in insert mode
au BufEnter * if &buftype == 'terminal' | :startinsert | endif
" open terminal on ctrl+n
function! OpenTerminal()
  split term://PowerShell
  resize 10
endfunction
nnoremap <c-n> :call OpenTerminal()<CR>


" use alt+hjkl to move between split/vsplit panels
tnoremap <A-h> <C-\><C-n><C-w>h
tnoremap <A-j> <C-\><C-n><C-w>j
tnoremap <A-k> <C-\><C-n><C-w>k
tnoremap <A-l> <C-\><C-n><C-w>l
nnoremap <A-h> <C-w>h
nnoremap <A-j> <C-w>j
nnoremap <A-k> <C-w>k
nnoremap <A-l> <C-w>l
nnoremap <Home> i
imap <Home> <ESC>

" file searching config
nnoremap <C-p> :FZF<CR>
let g:fzf_action = {
  \ 'ctrl-t': 'tab split',
  \ 'ctrl-s': 'split',
  \ 'ctrl-v': 'vsplit'
  \}

""""""""""""""""""""""""""""""""""""""""""
" MAPPINGS
""""""""""""""""""""""""""""""""""""""""""
inoremap jj <Esc>
nnoremap <leader>r :so ~/.vimrc<CR>

" Toggle NERDTree
nmap <leader>n :NERDTreeToggle<CR> 
nnoremap <silent><leader>/ :nohlsearch<Bar>:echo<CR>

" Use one of the following to define the camel characters.
let g:camelchar = "A-Z0-_9.,;:{([`'\""

nnoremap <silent><C-h> :<C-u>call search('\C\<\<Bar>\%(^\<Bar>[^'.g:camelchar.']\@<=\)['.g:camelchar.']\<Bar>['.g:camelchar.']\ze\%([^'.g:camelchar.']\&\>\@!\)\<Bar>\%^','bW')<CR>
nnoremap <silent><C-l> :<C-u>call search('\C\<\<Bar>\%(^\<Bar>[^'.g:camelchar.']\@<=\)['.g:camelchar.']\<Bar>['.g:camelchar.']\ze\%([^'.g:camelchar.']\&\>\@!\)\<Bar>\%$','W')<CR>
inoremap <silent><C-h> <C-o>:call search('\C\<\<Bar>\%(^\<Bar>[^'.g:camelchar.']\@<=\)['.g:camelchar.']\<Bar>['.g:camelchar.']\ze\%([^'.g:camelchar.']\&\>\@!\)\<Bar>\%^','bW')<CR>
inoremap <silent><C-l> <C-o>:call search('\C\<\<Bar>\%(^\<Bar>[^'.g:camelchar.']\@<=\)['.g:camelchar.']\<Bar>['.g:camelchar.']\ze\%([^'.g:camelchar.']\&\>\@!\)\<Bar>\%$','W')<CR>
vnoremap <silent><C-h> :<C-U>call search('\C\<\<Bar>\%(^\<Bar>[^'.g:camelchar.']\@<=\)['.g:camelchar.']\<Bar>['.g:camelchar.']\ze\%([^'.g:camelchar.']\&\>\@!\)\<Bar>\%^','bW')<CR>v`>o
vnoremap <silent><C-r> <Esc>`>:<C-U>call search('\C\<\<Bar>\%(^\<Bar>[^'.g:camelchar.']\@<=\)['.g:camelchar.']\<Bar>['.g:camelchar.']\ze\%([^'.g:camelchar.']\&\>\@!\)\<Bar>\%$','W')<CR>v`<o

" remove trailing white space
nnoremap <F5> :let _s=@/<Bar>:%s/\s\+$//e<Bar>:let @/=_s<Bar><CR>

" prevent search jump
\noremap * *``
" cycle between buffers, quickfix list and location list
nnoremap [n :bprev<cr>
nnoremap ]n :bnext<cr>
nnoremap [o :cprev<cr>
onoremap [o :cprev<cr>
nnoremap ]o :cnext<cr>
onoremap ]o :cnext<cr>
nnoremap -o :copen<cr>
nnoremap [l :lprev<cr>
onoremap [l :lprev<cr>
nnoremap ]l :lnext<cr>
onoremap ]l :lnext<cr>
nnoremap -l :lopen<cr>
nnoremap g! :tabedit <cfile><cr>
vnoremap g! y:tabedit <c-r>"<cr>
nnoremap gb :Git blame<cr><c-w>12<<cr>