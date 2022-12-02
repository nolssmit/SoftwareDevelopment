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
      body: Column(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: <Widget> [
          Row(
            children: const <Widget> [
              Text('Hello'),
              Text('world'),
            ],
          ),
          Container(
            padding: const EdgeInsets.all(20),
            color: Colors.cyan,
            child: const Text('One'),
          ),
                    Container(
            padding: const EdgeInsets.all(30),
            color: Colors.pinkAccent,
            child: const Text('Two'),
          ),
                    Container(
            padding: const EdgeInsets.all(40),
            color: Colors.amber,
            child: const Text('Three'),
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
