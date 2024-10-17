package com.paysystem.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.Type;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import java.time.LocalDateTime;
import java.util.List;

@Entity @NoArgsConstructor
@AllArgsConstructor
@Builder
public class Client {
    @Id @GeneratedValue(strategy = GenerationType.AUTO) @Type(type = "long")
    private Long id;
    private String name;
}
