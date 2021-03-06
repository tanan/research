import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
	id("org.springframework.boot")
	id("io.spring.dependency-management")
	kotlin("jvm")
	kotlin("plugin.spring")
	java
	application
}

group = "com.tanan.researchserver"
version = "0.0.1-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_11

dependencies {
	implementation("org.springframework.boot:spring-boot-starter")
	implementation("org.springframework.boot:spring-boot-starter-data-jpa")
	implementation("org.jetbrains.kotlin:kotlin-reflect")
	implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
	implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
	implementation("javax.inject:javax.inject:1")
	implementation(project(":research-server-domain"))
	implementation(project(":research-server-driver"))
	implementation(project(":research-server-port"))
}

tasks.withType<KotlinCompile> {
	kotlinOptions {
		freeCompilerArgs = listOf("-Xjsr305=strict")
		jvmTarget = "11"
	}
}

tasks.withType<org.springframework.boot.gradle.tasks.bundling.BootJar> {
	 manifest {
		 attributes["Main-Class"] = "com.tanan.researchserver.ResearchServerRestApplication"
	 }
	 val version = "1.0"
	 archiveName = "research-server.jar"
}