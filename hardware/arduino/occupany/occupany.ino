#include <LiquidCrystal.h>

// First we'll set up constants for the pin numbers.
// This will make it easier to follow the code below.

const int button1Pin = 2;  // pushbutton 1 pin
const int button2Pin = 3;  // pushbutton 2 pin
const int ledPin =  13;    // LED pin
String inputString = "";
boolean stringComplete = false;

LiquidCrystal lcd(12,11,7,6,5,4);

void setup()
{
  // Setup LCD
  lcd.begin(16, 2);
  lcd.clear();

  // Set up the pushbutton pins to be an input:
  pinMode(button1Pin, INPUT);
  pinMode(button2Pin, INPUT);

  // Set up the LED pin to be an output:
  pinMode(ledPin, OUTPUT);

  // Setup write to port
  Serial.begin(9600);
}

void loop()
{
  int button1State, button2State;  // variables to hold the pushbutton states

  button1State = digitalRead(button1Pin);
  button2State = digitalRead(button2Pin);

  if (stringComplete)
  {
    lcd.clear();
    lcd.print("RESERVED BY ");
    lcd.setCursor(0,1);
    lcd.print(inputString);
    inputString = "";
    stringComplete = false;
  }

  if (button1State == LOW)
  {
    digitalWrite(ledPin, HIGH);
    Serial.println("OCCUPIED");
    lcd.clear();
    delay(250);
  }

  if (button2State == LOW)
  {
    digitalWrite(ledPin, LOW);
    Serial.println("VACANT");
    delay(250);
  }

}

void serialEvent()
{
  while (Serial.available())
  {
    char inChar = (char)Serial.read();
    if (inChar == '!')
    {
      stringComplete = true;
    }
    inputString += inChar;
  }
}
