# What is it?
A tool inspired by many dotfile bootstrap  tools like (dotbot). Still work in progress.

# Why?
Just wanted a single binary that I can use to bootstrap my whole dotfiles without the need of installing anything.

# Sample Configuration

How to setup TOML configurations
```toml
[configs.<config_name>]
source = <source-of-the-config>
destination = <destination-of-the-config>
description = <description-what-is-the-config-is>

```
NOTES: 
When giving a destination of a file, the last part of the path should be the filename.
When specifying a directory to copy, the last folder specied is the folder where the files are going to be copied to.

Sample configuration
```toml
[configs.zsh]
source = ./.zshrc
destination = ~/.zshrc
description = "Copy zshrc config to home"

[configs.nvim]
source = ./nvim
destination = ~/.config/nvim
description = "Copy nvim configs to .config directory"
```

# License
Copyright 2024 norman santiago

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
