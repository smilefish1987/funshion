makedown就是一门标记语言

设计哲学   Markdown 的目标是易读易写。

Markdown强调可读性高于一切。一份Markdown格式的文档应该能直接以纯文本方式发布，而不致一眼看过去满眼都是标签和格式化指令。Markdown的语法确实受了几种现有的text转HTML过滤器影响－－包括 Setext, atx, Textile, reStructuredText,Grutatext, 和 EtText -- 其中对Markdown语法影响最大的单一来源是纯文本的Email格式。

为实现这一目标，Markdown的语法几乎全部由标点符号构成，这些标点符号都是精心挑选而来，尽量做到能望文生义。如星号括着一个单词（Markdown中表示强调）看上去就像 *强调*。Markdown的列表看上去就像列表；Markdown的引文就象引文，和你使用email时的感觉一样

这个一个普通段落

<table>
	<tr>
		<td> smilefish   *强调* </td>
	</tr>
</table>

这是另一个普通段落

&copyright

AT&T

4<5

This is an H1
=============

This is an H2
-------------

# This is another H1 #

## This is another H2

### This is h3 ###

#### This is h4

##### This is h5

###### This is h6




> This is a blockquote with two paragraphs. Lorem ipsum dolor sit amet,
> consectetuer adipiscing elit. Aliquam hendrerit mi posuere lectus.
> Vestibulum enim wisi, viverra nec, fringilla in, laoreet vitae, risus.
> 
> Donec sit amet nisl. Aliquam semper ipsum sit amet velit. Suspendisse
> id sem consectetuer libero luctus adipiscing.



> This is a blockquote with two paragraphs. Lorem ipsum dolor sit amet,
consectetuer adipiscing elit. Aliquam hendrerit mi posuere lectus.
Vestibulum enim wisi, viverra nec, fringilla in, laoreet vitae, risus.

> Donec sit amet nisl. Aliquam semper ipsum sit amet velit. Suspendisse
id sem consectetuer libero luctus adipiscing.


> This is the first level of quoting.
>
> > This is nested blockquote.
>
> Back to the first level.



> ## This is a header.
> 
> 1.   This is the first list item.
> 2.   This is the second list item.
> 
> Here's some example code:
> 
>     return shell_exec("echo $input | $markdown_script");



*   Red
*   Green
*   Blue

+   Red
+   Green
+   Blue

-   Red
-   Green
-   Blue


1.  Bird
2.  McHale
3.  Parish


*   Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
    Aliquam hendrerit mi posuere lectus. Vestibulum enim wisi,
    viverra nec, fringilla in, laoreet vitae, risus.
*   Donec sit amet nisl. Aliquam semper ipsum sit amet velit.
    Suspendisse id sem consectetuer libero luctus adipiscing.


*   Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
Aliquam hendrerit mi posuere lectus. Vestibulum enim wisi,
viverra nec, fringilla in, laoreet vitae, risus.
*   Donec sit amet nisl. Aliquam semper ipsum sit amet velit.
Suspendisse id sem consectetuer libero luctus adipiscing.



1.  This is a list item with two paragraphs. Lorem ipsum dolor
    sit amet, consectetuer adipiscing elit. Aliquam hendrerit
    mi posuere lectus.

    Vestibulum enim wisi, viverra nec, fringilla in, laoreet
    vitae, risus. Donec sit amet nisl. Aliquam semper ipsum
    sit amet velit.

2.  Suspendisse id sem consectetuer libero luctus adipiscing.



*   A list item with a blockquote:

    > This is a blockquote
    > inside a list item.



*   A list item with a code block:

        <code goes here>




This is a normal paragraph:

    This is a code block.



Here is an example of AppleScript:

    tell application "Foo"
        beep
    end tell




    <div class="footer">
        &copy; 2004 Foo Corporation
    </div>




*  *  *
hehehe
***

*****

-  -  -

------------------------------------