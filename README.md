# MarkMath 
This project wants to solve the problem of rendering math expressions in html. 
In html, you can use Mathjax to render math expressions but it's too heavy. MarkMath is a lighter markup language to write math expressions. 
MarkMath processor compiles the MarkMath texts into html. The compiled html files don't need Mathjax. They use svg to render so they're more independent. 
### Use 
Add the compiled executable application to environment paths and use ' MarkMath (filename) ' to compile MarkMath texts. The compiled result will be immediately printed to your screen. 
If you don't tell MarkMath which file you want to compile, it will start default edit mode. You can type your MarkMath texts and use ' \rend ' to compile and ' \quit ' to exit. 
### Grammar 
* Use ' $[...] ' to insert a  MarkMath expression in normal texts. 
* / => divide symbol 
* \* => multiply symbol 
* /[..., ...] => a/b 
### Will support 
* float nunber 
* root 
* triangle functions 
* power 
* below mark 