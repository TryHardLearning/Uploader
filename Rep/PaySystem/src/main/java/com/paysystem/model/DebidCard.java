package com.paysystem.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.Type;

import javax.persistence.*;
import java.math.BigInteger;

@Entity @NoArgsConstructor @AllArgsConstructor @Builder
public class DebidCard {
    @Id @GeneratedValue(strategy = GenerationType.AUTO) @Type(type = "long") @Getter
    private Long id;
    @Getter
    private BigInteger number;
    @Getter
    private Integer cvv;
    @ManyToOne
    private Client client;
}
