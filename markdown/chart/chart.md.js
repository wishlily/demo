(function($){
	var methods = {
		bar : function() {
			var data = $.parseJSON(this.text())
			// get data
			var datasets = [];
			var color = Math.random() * 360;
			for (var i = 0; i < data[`data`].length; i++) {
				var dataset = {};
				dataset.label = data[`data`][i][`label`];
				dataset.data = data[`data`][i][`data`];
				dataset.borderWidth = 1;
				dataset.hoverBorderWidth = 2;
				var h = Math.round(color + (i * 137.5)) % 360;
				var l = Math.round(Math.random() * 20) + 60;
				dataset.backgroundColor = 'hsla(' + h + ',70%,' + l + '%,0.4)';
				dataset.borderColor = 'hsla(' + h + ',70%,' + l + '%,1)';
				datasets[i] = dataset;
			}
			// get bar type
			var type = 'bar';
			var xstacked = false;
			var ystacked = false;
			if (data.hasOwnProperty("type")) {
				var ss = data[`type`];
				if (ss.indexOf('horiz') >= 0) {
					type = 'horizontalBar';
					if (ss.indexOf('over') >= 0) {
						ystacked = true;
					}
				} else {
					if (ss.indexOf('over') >= 0) {
						xstacked = true;
					}
				}
				if (ss.indexOf('stack') >= 0) {
					xstacked = true;
					ystacked = true;
				}
			}
			new Chart(this, {
				type: type,
				data: {
					labels: data[`item`],
					datasets: datasets,
				},
				options: {
					responsive: false,
					title: {
						display: data.hasOwnProperty("title"),
						text: data[`title`],
					},
					legend: {
						display: data[`data`][0].hasOwnProperty("label"),
					},
					scales: {
						xAxes: [{
							stacked: xstacked,
							ticks: {
								beginAtZero: true,
							},
						}],
						yAxes: [{
							stacked: ystacked,
							ticks: {
								beginAtZero: true,
							},
						}],
					},
					tooltips: {
						mode: (xstacked || ystacked) ? 'label' : 'single',
					},
				},
			});
		}
	};
	$.fn.chart = function( method ){
		if ( methods[method] ) {
			return methods[ method ].apply( this, Array.prototype.slice.call( arguments, 1 ));
		}  else if ( typeof method === 'object' || ! method ) {
			return methods.init.apply( this, arguments );
		} else {
			$.error( 'Method ' +  method + ' does not exist on jQuery.chart' );
		}
	};
})(jQuery);
