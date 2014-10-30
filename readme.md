#Tag Remover

This little utility is used to match and delete specific tags and all the children from HTML pages.

##Deprecated
Consider using my other library [nodefinder](https://github.com/meoow/nodefinder) for better path parsing and node finding (thus can do removing).  

##Usage

```sh
#rewite to stdout
tagremover webpage.html html/body/table/tbody[border=1,cellpadding=6]

#make changes inplace
#class="header"
tagremover -i webpage.html html/body/div.header

#examples go on

# use quotes or not treated as the same
tagremover webpage.html html/body/div.footer[id="foot footer", name=go python perl]

# match the second occurations
tagremove webpage.html html/body/2*div

```
