char cmd = 0;

void setup() {
    // Enable serial
    Serial.begin(9600);
    while(!Serial);
    // Enable LED
    pinMode(LED_BUILTIN, OUTPUT);
}

void loop() {
    // contrl LED blink
    switch (cmd) {
    case 'T':
        digitalWrite(LED_BUILTIN, HIGH);
        Serial.print("LED is on.\t[");
        Serial.print(digitalRead(LED_BUILTIN));
        Serial.println("]");
        cmd = 0;
        break;
    case 'F':
        digitalWrite(LED_BUILTIN, LOW);
        Serial.print("LED is off.\t[");
        Serial.print(digitalRead(LED_BUILTIN));
        Serial.println("]");
        cmd = 0;
        break;
    case 'L':
        Serial.println("GPIO Read List: ");
        for (int i = 0; i < 21; i++) {
            Serial.print(i);
            Serial.print("\t");
            Serial.print(digitalRead(i));
            Serial.println("");
        }
        cmd = 0;
        break;
    default:
        break;
    }
}

void serialEvent(){
    if (Serial.available() > 0) {
        cmd = Serial.read();
    }
}
