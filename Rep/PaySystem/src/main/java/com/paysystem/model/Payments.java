package com.paysystem.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.Type;

import javax.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.UUID;

@Entity @AllArgsConstructor @NoArgsConstructor
public class Payments {

    @Id @GeneratedValue(strategy = GenerationType.AUTO) @Type(type = "long")
    private Long id ;
    @Getter
    private BigDecimal price;
    @Getter
    private LocalDateTime date;
    @OneToOne @Getter
    private Client client;
    @Getter
    private String status;
    @OneToOne
    private Demand demand;
}
