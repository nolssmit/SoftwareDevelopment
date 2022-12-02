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

      // body: Center(
      // child: const Icon(
      //   Icons.airport_shuttle,
      //   color: Colors.lightBlue,
      //   size: 50,
      // ),

      // body: Center(
      //   child: RaisedButton(
      //     onPressed: () {},
      //     color: Color.fromARGB(255, 100, 175, 236),
      //     child: const Text(
      //       'Click me!',
      //     ),
      //   ),
      // ),

      // body: Center(
      //   child: RaisedButton.icon(
      //     onPressed: () {},
      //     icon: const Icon(
      //       color: Colors.blue,
      //       Icons.mail,
      //     ),
      //     label: Text('mail me'),
      //     color: Colors.amber,
      //   ),
      // ),

      body: Center(
        child: IconButton(
          onPressed: () {
            print('you clicked me!');
          },
          icon: const Icon(Icons.alternate_email),
          color: Colors.amber,
        ),
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
