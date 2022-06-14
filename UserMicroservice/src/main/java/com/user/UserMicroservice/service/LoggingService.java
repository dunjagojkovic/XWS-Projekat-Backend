package com.user.UserMicroservice.service;

import java.io.BufferedWriter;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.FileWriter;
import java.io.IOException;
import java.time.LocalDateTime;
import java.util.zip.ZipEntry;
import java.util.zip.ZipOutputStream;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import com.user.UserMicroservice.enums.LogEntryType;

@Service
public class LoggingService {
	@Value("${logging.enabled}")
	private boolean enabled;
	
	private String dataSeparator = " | ";
	
	public void log(LogEntryType type, String code, String ip, String user) throws IOException {
		boolean logLimitReached = false;
	    BufferedWriter writer = new BufferedWriter(new FileWriter(getFormattedFileName(type), !logLimitReached));
	    String logEntry = "[" + LocalDateTime.now() + "] " + code + dataSeparator + ip + dataSeparator + user + " ;";
	    writer.append(logEntry);
	    writer.newLine();
	    writer.close();
	    
	    File file = new File(getFormattedFileName(type));
	    if(file.exists() && file.length() >= 4000) {
	    	archiveLog(type.toString().toLowerCase());
	    	file.delete();
	    }
	}
	
	public void log(LogEntryType type, String code, String ip) throws IOException {
		log(type, code, ip, "*");
	}
	
	private void archiveLog(String fileName) throws IOException {
        String sourceFile = getFormattedFileName(fileName);
        FileOutputStream fos = new FileOutputStream("log//" + fileName + "_" + LocalDateTime.now().toString().replace(":", "_").replace(" ", "_") + ".zip");
        ZipOutputStream zipOut = new ZipOutputStream(fos);
        File fileToZip = new File(sourceFile);
        FileInputStream fis = new FileInputStream(fileToZip);
        ZipEntry zipEntry = new ZipEntry(fileToZip.getName());
        zipOut.putNextEntry(zipEntry);
        byte[] bytes = new byte[1024];
        int length;
        while((length = fis.read(bytes)) >= 0) {
            zipOut.write(bytes, 0, length);
        }
        zipOut.close();
        fis.close();
        fos.close();
    }
	
	private String getFormattedFileName(String fileName) {
		return "log//" + fileName + ".log";
	}
	
	private String getFormattedFileName(LogEntryType type) {
		return "log//" + type.toString().toLowerCase() + ".log";
	}
}
