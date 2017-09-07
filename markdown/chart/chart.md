<script src="/public/jquery-2.1.3.min.js"></script>
<script src="/public/Chart.min.js"></script>
<script src="/public/chart.md.js"></script>

sublime 使用 OmniMarkupPreviewer
复制 js 文件到 `xxx\AppData\Roaming\Sublime Text 3\Packages\OmniMarkupPreviewer\public` 目录下

表格测试1：

<canvas class="bar" width="400" height="400">
{
	"title": "I have Title",
	"item": ["Red", "Blue", "Yellow", "Green", "Purple", "Orange"],
	"type": "horiz over stack",
	"data": [{
			"label": "one",
			"data": [12, 19, 3, 5, 2, 3]
		},{
			"label": "two",
			"data": [23, 34, 1, 0, 39, 22]
		},{
			"label": "thrid",
			"data": [21, 0, 1, 20, 12, 15]
		}]
}
</canvas>

表格测试2：

<canvas class="bar" width="400" height="400">
{
	"item": ["Red", "Blue", "Yellow", "Green", "Purple", "Orange"],
	"data": [{
			"data": [33, 19, 3, 5, 2, 3]
		}]
}
</canvas>

<script>
	$(".bar").each(function(){
		$(this).chart('bar')
	})
</script>
