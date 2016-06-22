import 'file?name=[name].[ext]!./index.html'
import React from 'react'
import ReactDom from 'react-dom'

class Hello extends React.Component {
	render(){
	return <div>
			<h1>{"Hello World!"}</h1>
		</div>
	}
}

ReactDom.render(
	<Hello />,
	document.getElementById('example'));
