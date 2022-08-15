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
      //   padding: EdgeInsets.symmetric(horizontal: 30, vertical: 10),
      //   margin: EdgeInsets.all(30),
      //   color: Colors.grey[400],
      //   child: Text('Hello'),
      // ),

      body: Padding(
        padding: EdgeInsets.all(90),
        child: Text('Hello'),
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
