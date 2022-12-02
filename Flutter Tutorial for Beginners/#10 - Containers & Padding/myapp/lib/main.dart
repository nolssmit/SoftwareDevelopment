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
      
      // body: Container(
      //   padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 10),
      //   margin: const EdgeInsets.all(30),
      //   color: Colors.grey[400],
      //   child: const Text('Hello'),
      // ),

      body: const Padding(
        padding: EdgeInsets.all(90),
        child: Text('Hello'),
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
