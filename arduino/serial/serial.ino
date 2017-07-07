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
        Serial.printf("LED is on [%d].\n", digitalRead(LED_BUILTIN));
        cmd = 0;
        break;
    case 'F':
        digitalWrite(LED_BUILTIN, LOW);
        Serial.printf("LED is off [%d].\n", digitalRead(LED_BUILTIN));
        cmd = 0;
        break;
    case 'L':
        Serial.println("GPIO Read List: ");
        for (int i = 0; i < 21; i++) {
            Serial.printf("%d\t%d\n", i, digitalRead(i));
        }
        cmd = 0;
        break;
    default:
        break;
    }
}

void serialEvent(){
    cmd = Serial.read();
}
