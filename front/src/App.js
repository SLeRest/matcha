import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
	return (
		<FlatList
		data={DATA}
		renderItem={({ item }) => <Item title={item.title} />}
		keyExtractor={item => item.id}
		/>
	)
}

export default App;
