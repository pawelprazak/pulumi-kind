plugins {
    application
}

repositories {
    mavenLocal()
    mavenCentral()
}

dependencies {
    implementation("io.pulumi:pulumi:3.6.0+")
    implementation("io.pulumi:kubernetes:3.6.0+")
    implementation("com.google.code.findbugs:jsr305:3.0.2")
}

application {
    mainClass.set("org.example.Main")
}

run {
    logging.captureStandardError(LogLevel.INFO)
    logging.captureStandardOutput(LogLevel.INFO)
}

