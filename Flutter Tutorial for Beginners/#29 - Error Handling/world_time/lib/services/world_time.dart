import 'package:http/http.dart';
import 'dart:convert';

class WorldTime {
  String location; // location name for the UI
  late String time; // the time in that location
  String flag; // url to an asset flag icon
  String url; // location url for api endpoint

  WorldTime({required this.location, required this.flag, required this.url});

  Future<void> getTime() async {
    try {
      // Make the request
      Response response = await get(
        Uri.parse('http://worldtimeapi.org/api/timezone/$url'),
      );
      Map data = jsonDecode(response.body);
      //print(data);

      // get properties from data
      String datetime = data['datetime'];
      //print('datetime $datetime');
      //print('utc_offset: ${data['utc_offset']}');
      String offset = data['utc_offset'].substring(1, 3);
      //print('offset: $offset');

      // create DateTime object
      DateTime now = DateTime.parse(datetime);
      //print('now without offset: $now');
      now = now.add(Duration(hours: int.parse(offset)));
      //print('now with offset: $now');

      // set the time  property
      time = now.toString();
    } catch (e) {
      print('caught error $e');
      time = 'could not get time data';
    }
  }
}
