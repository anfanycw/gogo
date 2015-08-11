const int occupiedPin = 2;  // pushbutton 1 pin
const int vacantPin = 3;  // pushbutton 2 pin
const int ledPin =  13;    // LED pin

void setup()
{
  // Set up the pushbutton pins to be an input:
  pinMode(occupiedPin, INPUT);
  pinMode(vacantPin, INPUT);

  // Set up the LED pin to be an output:
  pinMode(ledPin, OUTPUT);      
  
  // Setup write to port
  Serial.begin(9600);
}

void loop()
{
  int occupiedState, vacantState;  // variables to hold the pushbutton states

  occupiedState = digitalRead(occupiedPin);
  vacantState = digitalRead(vacantPin);

  if (occupiedState == LOW)
  {
    digitalWrite(ledPin, HIGH);
    Serial.println("OCCUPIED");
    delay(500);
  } 

  if (vacantState == LOW)
  {
    digitalWrite(ledPin, LOW);
    Serial.println("VACANT");
    delay(500);
  }
}
