package com.agent.service;

import com.agent.model.User;
import net.bytebuddy.utility.RandomString;
import org.aspectj.apache.bcel.classfile.Code;
import org.springframework.stereotype.Service;

@Service
public class CodeService {
    public String generateActivationCodeForUSer(User user) {
        return RandomString.make(64);
    }

    public String generatePasswordResetCode(User user) {
        return RandomString.make(64);
    }

    public String generateLoginCode(User user) {
        // TODO Auto-generated method stub
        return RandomString.make(8);
    }

}
