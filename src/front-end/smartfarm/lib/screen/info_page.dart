import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:smartfarm/components/body.dart';
import 'package:smartfarm/constants/smartfarmer_constants.dart';

class InfoPage extends StatefulWidget {
  @override
  _InfoPageState createState() => _InfoPageState();
}

class _InfoPageState extends State<InfoPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: BottomAppBar(
        //color: appBarColor,
        //notchMargin: 12.0,
        shape: CircularNotchedRectangle(),
        child: Container(
          height: 60.0,
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 42.0),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: <Widget>[
                Icon(Icons.home, color: blueGradient1, size: 32.0),
                Icon(Icons.settings, color: Colors.grey, size: 32.0),
              ],
            ),
          ),
        ),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerDocked,
      floatingActionButton: FloatingActionButton(
        child: Icon(Icons.add),
        onPressed: (){},
      ),
      appBar: buildAppBar(),
      backgroundColor: bgColor,
      body: Body(),
    );
  }

  AppBar buildAppBar() {
    return AppBar(
      elevation: 0,
      flexibleSpace: Container(
        decoration: BoxDecoration(
          gradient: LinearGradient(
            colors: [blueGradient1, blueGradient2],
            begin: Alignment.bottomLeft,
            end: Alignment.bottomRight,
          ),
        ),
      ),
//      elevation: 0,
      leading: IconButton(
        icon: SvgPicture.asset("assets/icons/menu.svg"),
        onPressed: () {},
      ),
      actions: <Widget>[
        IconButton(
          icon: Icon(Icons.notifications_none),
          color: Colors.white,
          onPressed: () {},
        )
      ],
    );
  }
}
