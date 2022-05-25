package com.user.UserMicroservice.config;

import com.user.UserMicroservice.security.TokenAuthenticationFilter;
import com.user.UserMicroservice.security.TokenUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.builders.WebSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

@Configuration
@EnableWebSecurity
@EnableGlobalMethodSecurity(prePostEnabled = true)
public class WebSecurityConfig extends WebSecurityConfigurerAdapter {

    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }

    @Autowired
    private CustomUserDetailsService jwtUserDetailsService;

    @Bean
    @Override
    public AuthenticationManager authenticationManagerBean() throws Exception {
        return super.authenticationManagerBean();
    }

    @Autowired
    public void configureGlobal(AuthenticationManagerBuilder auth) throws Exception {
        auth.userDetailsService(jwtUserDetailsService).passwordEncoder(passwordEncoder());
    }

    @Autowired
    private TokenUtil tokenUtils;

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS).and()

                .authorizeRequests().antMatchers("/api/users/register").permitAll()
                .antMatchers("/error").permitAll()
                .antMatchers("/error/**").permitAll()
                .antMatchers("/your Urls that dosen't need security/**").permitAll()
                .antMatchers("/api/users/forgottenpassword").permitAll()
                .antMatchers("/api/users/checkActivationCode").permitAll()
                .antMatchers("/api/users/checkForgottenPassword").permitAll()
                .antMatchers("/api/users/loginCode").permitAll()

                


                .anyRequest().authenticated().and()

                //.formLogin().loginPage("/login").permitAll().and()

                .logout().permitAll().and()

                .cors().and()

                .httpBasic().and()

                .addFilterBefore(new TokenAuthenticationFilter(tokenUtils, jwtUserDetailsService), BasicAuthenticationFilter.class);
        http.csrf().disable();
        
        http
        .headers()
        .xssProtection()
        .and()
        .contentSecurityPolicy("script-src 'self'");
    }

    //OVDE SE STAVLJA ONO ZA STA NE TREBA TOKEN, NE DODAVATI SVE MOGUCE RUTE!!!!!!!!!!!!!
    @Override
    public void configure(WebSecurity web) throws Exception {
        web.ignoring().antMatchers(HttpMethod.POST, "/api/users/login", "api/users/register");
        web.ignoring().antMatchers(HttpMethod.OPTIONS, "/**");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/follow/follower");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/follow/accept");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/follow/follower");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/users/checkActivationCode");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/users/checkForgottenPassword");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/users/forgottenpassword");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/users/filterUsers");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/users/loginCode");
        web.ignoring().antMatchers(HttpMethod.GET, "/api/users/public");
        web.ignoring().antMatchers(HttpMethod.GET, "/api/follow/following/**");
        web.ignoring().antMatchers("/error/**");
    }
}