<script src="/public/jquery-2.1.3.min.js"></script>
<script src="/public/underscore-min.js"></script>
<script src="/public/raphael.min.js"></script>
<script src="/public/sequence-diagram-min.js"></script>

sublime 使用 OmniMarkupPreviewer
复制 js 文件到 `xxx\AppData\Roaming\Sublime Text 3\Packages\OmniMarkupPreviewer\public` 目录下

UML1:

<div class="diagram">
	Alice->Bob: Hello Bob, how are you?
	Note right of Bob: Bob thinks
	Bob-->Alice: I am good thanks!
</div>

UML2:

<div class="diagram">
	Title: Here is a title
	A->B: Normal line
	B-->C: Dashed line
	C->>D: Open arrow
	D-->>A: Dashed open arrow
</div>

<script>
	$(".diagram").sequenceDiagram({theme: 'hand'});
</script>