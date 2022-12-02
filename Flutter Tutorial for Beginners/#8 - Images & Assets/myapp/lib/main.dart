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
      body: Center(
//        child: Image.asset('images/a.jpg'),
        child: Image.network('https://images.unsplash.com/photo-1659221483823-b9dcb02dc6c3?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80'),
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
