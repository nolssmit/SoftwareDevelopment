import 'package:flutter/material.dart';

void main() => runApp(MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text(
            'My First App',
          ),
          centerTitle: true,
          backgroundColor: Colors.red[700],
        ),
        body: const Center(
          child: Text(
            'Hello daar!',
            style: TextStyle(
              fontSize: 20,
              fontWeight: FontWeight.bold,
              letterSpacing: 2,
              color: Colors.grey,
              fontFamily: 'IndieFlower',            
            ),
          ),
        ),
        floatingActionButton: const FloatingActionButton(
          onPressed: null,
          child: Text(
            'click',
          ),
          backgroundColor: Color.fromARGB(255, 248, 20, 4),
        ),
      ),
    ));
