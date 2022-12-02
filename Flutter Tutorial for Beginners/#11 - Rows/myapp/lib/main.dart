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
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          const Text('Hello world!'),
          FlatButton(
            onPressed: () {},
            color: Colors.amber,
            child: const Text('click me'),
          ),
          Container(
            color: Colors.cyan,
            padding: const EdgeInsets.all(30),
            child: const Text('inside container'),
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
