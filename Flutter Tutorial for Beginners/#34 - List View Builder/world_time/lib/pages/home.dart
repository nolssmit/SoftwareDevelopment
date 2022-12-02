import 'package:flutter/material.dart';
import 'dart:convert';

class Home extends StatefulWidget {
  @override
  _HomeState createState() => _HomeState();
}

class _HomeState extends State<Home> {
  Map data = {};
  Object? parameters;

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    parameters = ModalRoute.of(context)!.settings.arguments;
    Map data = jsonDecode(jsonEncode(parameters));
    print(data);

    // set background
    String bgImage = data['isDaytime'] ? 'day.png' : 'night.png';
    Color bgColor =
        data['isDaytime'] ? Colors.lightBlue : Colors.indigo[700] as Color;

    return Scaffold(
      backgroundColor: bgColor,
      //appBar: AppBar(),
      body: SafeArea(
        child: Container(
          decoration: BoxDecoration(
            image: DecorationImage(
              image: AssetImage('assets/$bgImage'),
              fit: BoxFit.cover,
            ),
          ),
          child: Padding(
            padding: const EdgeInsets.fromLTRB(0, 120.0, 0, 0),
            child: Column(
              children: <Widget>[
                FlatButton.icon(
                  onPressed: () {
                    Navigator.pushNamed(context, '/location');
                  },
                  icon: const Icon(
                    Icons.edit_location,
                    color: Color.fromARGB(255, 221, 218, 218),
                  ),
                  label: const Text(
                    'Edit Location',
                    style: TextStyle(color: Colors.grey),
                  ),
                ),
                const SizedBox(height: 20.0),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    Text(
                      data['location'],
                      style: const TextStyle(
                        fontSize: 28.0,
                        letterSpacing: 2.0,
                        color: Colors.white,
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: 20.0),
                Text(data['time'],
                    style: const TextStyle(
                      fontSize: 66.0,
                      color: Colors.white,
                    ),),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
