# GORM Gen 사용해보기 


## 1. GORM Gen 이란?
GEN: 코드 생성으로 구동되는 친근하고 안전한 GORM.

### 개요
- 동적 원시 SQL의 관용적이고 재사용 가능한 API
- 인터페이스가 없는 100% 유형 안전 DAO API{}
- 데이터베이스 투 스트럭처는 GORM 규칙을 따릅니다.
- 내부적으로 GORM이 지원하는 모든 기능, 플러그인, DBMS 지원

### GORM Gen의 장점
- Query 에 대한 직관적인 이해가 쉬움
- DB 구조 및 Query 의 변경이 Code화 되어 런타임전에 오류 검증이 가능

### GORM Gen의 단점
- Sql 생성시점에 DB의 구조를 검증하지 않음 (DB의 구조가 변경되었을 경우, 런타임에 오류가 발생할 수 있음)

