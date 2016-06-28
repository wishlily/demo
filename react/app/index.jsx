import 'file?name=[name].[ext]!./index.html'
import React from 'react'
import ReactDom from 'react-dom'
import Brief from './extend.jsx'

class HelloWorld extends React.Component {
	render(){
	return <div>
			<h1>{"Hello World!"}</h1>
			<Brief />
		</div>
	}
}

ReactDom.render(
	<HelloWorld />,
	document.getElementById('example'));
