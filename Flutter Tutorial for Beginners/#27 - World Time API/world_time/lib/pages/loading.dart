// For the duumy data used, see:  https://jsonplaceholder.typicode.com/

import 'package:flutter/material.dart';
import 'package:http/http.dart';
import 'dart:convert';

class Loading extends StatefulWidget {
  @override
  State<Loading> createState() => _LoadingState();
}

class _LoadingState extends State<Loading> {
  int counter = 0;

  void getTime() async {
    // Make the request
    Response response = await get(
      Uri.parse('http://worldtimeapi.org/api/timezone/Africa/Johannesburg'),
    );
    Map data = jsonDecode(response.body);
    //print(data);

    // get properties from data
    String datetime = data['datetime'];
    print('datetime: $datetime');
    print('utc_offset: ${data['utc_offset']}');
    String offset = data['utc_offset'].substring(1, 3);
    print('offset: $offset');

    // create DateTime object
    DateTime now = DateTime.parse(datetime);
    print('now without offset: $now');
    now = now.add(Duration(hours: int.parse(offset)));
    print('now with offset: $now');
  }

  @override
  void initState() {
    super.initState();
    getTime();
  }

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Text('Loading screen'),
    );
  }
}
