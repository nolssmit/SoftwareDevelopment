// https://www.youtube.com/watch?v=zNZvuP8h1vs&list=PL4cUxeGkcC9jLYyp2Aoh6hcWuxFDX6PBJ&index=14

import 'package:flutter/material.dart';

void main() => runApp(MaterialApp(
      home: Home(),
    ));

class Home extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text(
          'My First App',
        ),
        centerTitle: true,
        backgroundColor: Colors.red[700],
      ),
      body: Row(
        children: <Widget>[
          Expanded(
            flex: 5,
            child: Image.asset('images/a.jpg'),
          ),
          Expanded(
            flex: 1,
            child: Container(
              padding: const EdgeInsets.all(30),
              color: Colors.cyan,
              child: const Text('1'),
            ),
          ),
          Expanded(
            flex: 1,
            child: Container(
              padding: const EdgeInsets.all(30),
              color: Colors.pinkAccent,
              child: const Text('2'),
            ),
          ),
          Expanded(
            flex: 1,
            child: Container(
              padding: const EdgeInsets.all(30),
              color: Colors.amber,
              child: const Text('3'),
            ),
          ),
        ],
      ),
      floatingActionButton: const FloatingActionButton(
        onPressed: null,
        backgroundColor: Color.fromARGB(255, 248, 20, 4),
        child: Text(
          'click',
        ),
      ),
    );
  }
}
