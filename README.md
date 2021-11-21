# MarkMath 
![](icon/MarkMathIcon.jpeg) 
This project wants to solve the problem of rendering math expressions in html. 
In html, you can use Mathjax to render math expressions but it's too heavy. MarkMath is a lighter markup language to write math expressions. 
MarkMath processor compiles the MarkMath texts into html. The compiled html files don't need Mathjax. They use svg to render so they're more independent. 
### Use 
Add the compiled executable application to environment paths and use ' $ MarkMath (filename) ' to compile MarkMath texts. The compiled result will be immediately printed to your screen. 
If you don't tell MarkMath which file you want to compile, it will start default edit mode. You can type your MarkMath texts and use ' \rend ' to compile and ' \quit ' to exit. 
### Grammar 
* Use ' $[...] ' to insert a  MarkMath expression in normal texts. 
* / => divide symbol 
* \* => multiply symbol 
* /[..., ...] => a/b 
* +- => (You know what I mean) 
* ;a => α
* ;b => β
* ;t => θ
* ;p => π
* ;g => γ
* ;l => λ
* ;d => ΔΑ
### Will support 
* root ( need debugging )
* triangle functions 
* power 
* below mark 
* better error & warning report