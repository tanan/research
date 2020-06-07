@file:Suppress("DEPRECATION")

import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
	id("org.springframework.boot") version "2.2.4.RELEASE"
	id("io.spring.dependency-management") version "1.0.9.RELEASE"
	id("com.palantir.docker") version "0.22.1"
	kotlin("jvm") version "1.3.61"
	kotlin("plugin.spring") version "1.3.61"
}

group = "com.tanan"
version = "0.0.1-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_11

allprojects {
	repositories {
		mavenCentral()
		jcenter()
	}
}

subprojects {
	apply(plugin = "kotlin")
	apply(plugin = "java")
	apply(plugin = "application")
}

//
//springBoot {
//	mainClassName = "com.tanan.researchserver"
//}

dependencies {
	implementation("org.jetbrains.kotlin:kotlin-reflect")
	implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
//	implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
}

//tasks.withType<Test> {
//	useJUnitPlatform()
//}
//
//tasks.withType<KotlinCompile> {
//	kotlinOptions {
//		freeCompilerArgs = listOf("-Xjsr305=strict")
//		jvmTarget = "1.8"
//	}
//}


//tasks.withType<Jar> {
//	manifest {
//		attributes["Main-Class"] = "com.tanan.researchserver"
//	}
//	val version = "1.0"
//	archiveName = "research-server.jar"
//}
//
//tasks {
//	docker {
//		name = "history-log"
//		files(jar.name)
//		setDockerfile(file("$projectDir/Dockerfile"))
//	}
//}