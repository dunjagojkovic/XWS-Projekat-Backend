package com.agent.service;

import com.agent.service.formatter.AccountActivationFormatter;
import com.agent.service.formatter.LoginCodeFormatter;
import com.agent.service.formatter.ResetPasswordFormatter;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

@Service
public class MailService<T> {
    private Environment env;
    private JavaMailSender mailSender;

    @Autowired
    public MailService(Environment env, JavaMailSender mailSender) {
        this.env = env;
        this.mailSender = mailSender;
    }

    @Async
    public void sendUserRegistrationMail(String recipient, String activationCode, String siteUrl) {
        SimpleMailMessage message = new SimpleMailMessage();
        AccountActivationFormatter formatter = new AccountActivationFormatter();
        message.setFrom("dislinkt.xml@gmail.com");
        message.setTo(recipient);
        message.setText(formatter.getText(activationCode, siteUrl));
        message.setSubject(formatter.getSubject());


        mailSender.send(message);
    }

    @Async
    public void sendLinkToResetPassword(String recipient, String code, String siteUrl) {
        SimpleMailMessage message = new SimpleMailMessage();
        ResetPasswordFormatter formatter = new ResetPasswordFormatter();
        message.setFrom("dislinkt.xml@gmail.com");
        message.setTo(recipient);
        message.setText(formatter.getText(code, siteUrl));
        message.setSubject(formatter.getSubject());

        mailSender.send(message);
    }

    @Async
    public void sendCodetToEmail(String email, String loginCode, String siteURL) {
        SimpleMailMessage message = new SimpleMailMessage();
        LoginCodeFormatter formatter = new LoginCodeFormatter();
        message.setFrom("dislinkt.xml@gmail.com");
        message.setTo(email);
        message.setText(formatter.getText(loginCode, siteURL));
        message.setSubject(formatter.getSubject());

        mailSender.send(message);
    }
}
