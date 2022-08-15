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
            children: <Widget> [
              Text('Hello'),
              Text('world'),
            ],
          ),
          Container(
            padding: EdgeInsets.all(20),
            color: Colors.cyan,
            child: Text('One'),
          ),
                    Container(
            padding: EdgeInsets.all(30),
            color: Colors.pinkAccent,
            child: Text('Two'),
          ),
                    Container(
            padding: EdgeInsets.all(40),
            color: Colors.amber,
            child: Text('Three'),
          ),
        ],
      ),
      floatingActionButton: const FloatingActionButton(
        onPressed: null,
        child: Text(
          'click',
        ),
        backgroundColor: Color.fromARGB(255, 248, 20, 4),
      ),
    );
  }
}
