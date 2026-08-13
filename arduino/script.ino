#define LED 2

const int delay_value = 500;
bool printed = false;

void setup() {
  Serial.begin(115200);

  pinMode(LED, OUTPUT);
}

void loop() {
  if (Serial.available() > 0) {
    String message = Serial.readStringUntil('\n');

    float memoryUsage = message.toFloat();

    if(memoryUsage > 70.0) {
        digitalWrite(LED, HIGH);
    }
  }
}