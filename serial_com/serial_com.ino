#include <SoftwareSerial.h>

int byteBuff = 0;
bool b[4];
int p[4];

const int P1 = 2;
const int P2 = 3;
const int P3 = 4;
const int P4 = 5;

void setup() {
  pinMode(P1, OUTPUT); 
  pinMode(P2, OUTPUT); 
  pinMode(P3, OUTPUT); 
  pinMode(P4, OUTPUT);
  p[0] = P1;
  p[1] = P2;
  p[2] = P3;
  p[3] = P4;
  
  // Open serial communications and wait for port to open:
  Serial.begin(9600);
  while (!Serial) {
    ; // wait for serial port to connect. Needed for native USB port only
  }


  Serial.println("Relay on duty!");
}

void loop() { // run over and over
  String message;
  int i;
  
  if (Serial.available()) {
    message = Serial.readStringUntil('#');
    byteBuff = Serial.read();
    if (message == "allon") {
      for (i = 0; i < 4; i++) {
        b[i] = false;
        digitalWrite(p[i], LOW);
      };
    }
    else if (message == "alloff") {
      for (i = 0; i < 4; i++) {
        b[i] = true;
        digitalWrite(p[i], HIGH);
      };
    }
    else if (message == "status") {
      String outputMessage = "";
      for (i = 0; i < 4; i++) {
        outputMessage+= b[i];
      };
      outputMessage += '#';
      Serial.println(outputMessage);
    }
    else if (message.startsWith("set")) {
      i = message[3] - '0';
      if (i < 1 || i > 4) {
        Serial.println("invalid index#");
      }
      i--;
      if (message.endsWith("off")) {
        b[i] = true;
        digitalWrite(p[i], HIGH);
      }
      else if (message.endsWith("on")) {
        b[i] = false;
        digitalWrite(p[i], LOW);
      }
    }
    else if (message.startsWith("toggle")) {
      i = message[6] - '0';
      if (i > 0 && i <= 4) {
        i--;
        b[i] = !b[i];
        if (b[i]) {
          digitalWrite(p[i], HIGH);
        }
        else {
          digitalWrite(p[i], LOW);
        }
      } else {
        Serial.println("invalid index#");
      }
    }
    else {
      Serial.println("message: '" + message + "' is invalid#");
    }
    Serial.readString();
    Serial.flush();
  } 

  //Serial.println();
  //Serial.println(b);
  //Serial.write(Serial.read());
  
}
